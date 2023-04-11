import {
  Box,
  Button,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  HStack,
  Icon,
  Input,
  Select,
  Text,
  useColorModeValue,
  VStack,
  Wrap,
  WrapItem,
} from "@chakra-ui/react"
import { useQuery } from "@tanstack/react-query"
import { useCallback, useEffect, useMemo, useState } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { FiFolder } from "react-icons/fi"
import { useNavigate } from "react-router"
import { useSearchParams } from "react-router-dom"
import { client } from "../../client"
import { CollapsibleSection, useStreamingTerminal } from "../../components"
import { useProviders, useWorkspace, useWorkspaces } from "../../contexts"
import {
  CppSvg,
  DotnetcorePng,
  GoPng,
  JavaPng,
  NodejsPng,
  PhpSvg,
  PythonSvg,
  RustSvg,
} from "../../images"
import { exists, useFormErrors } from "../../lib"
import { QueryKeys } from "../../queryKeys"
import { Routes } from "../../routes"
import { useBorderColor } from "../../Theme"
import { TProviderID, TWorkspaceID } from "../../types"
import { ExampleCard } from "./ExampleCard"
import { FieldName, TFormValues } from "./types"

const DEFAULT_PREBUILD_REPOSITORY_KEY = "CREATE_PREBUILD_REPOSITORY"

// TODO: handle no provider configured
export function CreateWorkspace() {
  const idesQuery = useQuery({
    queryKey: QueryKeys.IDES,
    queryFn: async () => (await client.ides.listAll()).unwrap(),
  })

  const [isSubmitLoading, setIsSubmitLoading] = useState(false)
  const params = useCreateWorkspaceParams()
  const [workspaceID, setWorkspaceID] = useState<TWorkspaceID | undefined>(undefined)
  const navigate = useNavigate()
  const workspaces = useWorkspaces()
  const workspace = useWorkspace(workspaceID)
  const [[providers]] = useProviders()
  const { register, handleSubmit, formState, watch, setError, setValue, clearErrors } =
    useForm<TFormValues>({
      defaultValues: {
        [FieldName.PREBUILD_REPOSITORY]:
          window.localStorage.getItem(DEFAULT_PREBUILD_REPOSITORY_KEY) || "",
      },
    })
  const currentSource = watch(FieldName.SOURCE)
  const { terminal, connectStream } = useStreamingTerminal()

  useEffect(() => {
    if (params.rawSource !== undefined) {
      setValue(FieldName.SOURCE, params.rawSource)
    }

    // default ide
    if (params.ide !== undefined) {
      setValue(FieldName.DEFAULT_IDE, params.ide)
    } else if (idesQuery.data?.length) {
      const defaultIDE = idesQuery.data.find((ide) => ide.default)
      if (defaultIDE) {
        setValue(FieldName.DEFAULT_IDE, defaultIDE.name!)
      } else {
        const vscode = idesQuery.data.find((ide) => ide.name === "vscode")
        if (vscode) {
          setValue(FieldName.DEFAULT_IDE, vscode.name!)
        }
      }
    }

    // default provider
    if (params.providerID !== undefined) {
      setValue(FieldName.PROVIDER, params.providerID)
    } else if (providers) {
      const defaultProviderID = Object.keys(providers).find(
        (providerID) => providers[providerID]?.default
      )
      if (defaultProviderID) {
        setValue(FieldName.PROVIDER, defaultProviderID)
      }
    }
  }, [params, idesQuery.data, providers, setValue])

  const onSubmit = useCallback<SubmitHandler<TFormValues>>(
    async (data) => {
      // save prebuild repository
      if (data[FieldName.PREBUILD_REPOSITORY]) {
        window.localStorage.setItem(
          DEFAULT_PREBUILD_REPOSITORY_KEY,
          data[FieldName.PREBUILD_REPOSITORY]
        )
      } else {
        window.localStorage.removeItem(DEFAULT_PREBUILD_REPOSITORY_KEY)
      }

      const workspaceSource = data[FieldName.SOURCE].trim()
      setIsSubmitLoading(true)
      let workspaceID = data[FieldName.ID]
      if (!workspaceID) {
        const newIDResult = await client.workspaces.newID(workspaceSource)
        if (newIDResult.err) {
          setIsSubmitLoading(false)
          setError(FieldName.SOURCE, { message: newIDResult.val.message })

          return
        }

        workspaceID = newIDResult.val
      }

      if (workspaces.find((workspace) => workspace.id === workspaceID)) {
        setIsSubmitLoading(false)
        setError(FieldName.SOURCE, { message: "workspace with the same name already exists" })

        return
      }

      const providerID = data[FieldName.PROVIDER]
      const defaultIDE = data[FieldName.DEFAULT_IDE]

      // set default provider
      const useProviderResult = await client.providers.useProvider(providerID)
      if (useProviderResult.err) {
        setIsSubmitLoading(false)
        setError(FieldName.SOURCE, { message: useProviderResult.val.message })

        return
      }

      // set default ide
      const useIDEResult = await client.ides.useIDE(defaultIDE)
      if (useIDEResult.err) {
        setIsSubmitLoading(false)
        setError(FieldName.SOURCE, { message: useIDEResult.val.message })

        return
      }

      // create workspace
      workspace.create(
        {
          id: workspaceID,
          prebuildRepositories: data[FieldName.PREBUILD_REPOSITORY]
            ? [data[FieldName.PREBUILD_REPOSITORY]]
            : [],
          providerConfig: { providerID },
          ideConfig: { name: defaultIDE },
          sourceConfig: {
            source: workspaceSource,
          },
        },
        connectStream
      )

      // set workspace id to show terminal
      setWorkspaceID(workspaceID)
    },
    [workspaces, workspace, connectStream, setError]
  )

  const { sourceError, providerError, defaultIDEError, idError, prebuildRepositoryError } =
    useFormErrors(Object.values(FieldName), formState)

  const providerOptions = useMemo<readonly TProviderID[]>(() => {
    if (!exists(providers)) {
      return [] // TODO: make dynamic
    }

    return Object.keys(providers)
  }, [providers])

  const isLoading = useMemo(() => workspace.current?.name === "create", [workspace])

  const handleSelectFolderClicked = useCallback(async () => {
    const selected = await client.selectFromDir()
    if (selected) {
      setValue(FieldName.SOURCE, selected + "", {
        shouldDirty: true,
      })
    }
  }, [setValue])

  useEffect(() => {
    if (
      workspace.current?.name === "create" &&
      workspace.current.status === "success" &&
      workspace.data?.id !== undefined
    ) {
      navigate(Routes.WORKSPACES)
    }
  }, [navigate, workspace])

  const backgroundColor = useColorModeValue("blackAlpha.100", "whiteAlpha.100")
  const borderColor = useBorderColor()
  const inputBackgroundColor = useColorModeValue("white", "black")

  if (isLoading) {
    return terminal
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <VStack align="start" spacing="6" marginBottom="8">
        <VStack
          width="full"
          backgroundColor={backgroundColor}
          borderRadius="lg"
          borderWidth="thin"
          borderColor={borderColor}>
          <FormControl
            padding="20"
            isRequired
            isInvalid={exists(sourceError)}
            justifyContent={"center"}
            borderBottomWidth="thin"
            borderBottomColor={borderColor}>
            <VStack>
              <Text marginBottom="2" fontWeight="bold">
                Enter any git repository or local path to a folder you would like to create a
                workspace from
              </Text>
              <HStack spacing={0} justifyContent={"center"}>
                <Input
                  backgroundColor={inputBackgroundColor}
                  borderTopRightRadius={0}
                  borderBottomRightRadius={0}
                  placeholder="github.com/my-org/my-repo"
                  fontSize={"16px"}
                  padding={"10px"}
                  height={"42px"}
                  width={"400px"}
                  type="text"
                  {...register(FieldName.SOURCE, { required: true })}
                />
                <Button
                  leftIcon={<Icon as={FiFolder} />}
                  borderTopLeftRadius={0}
                  borderBottomLeftRadius={0}
                  borderTop={"1px solid white"}
                  borderRight={"1px solid white"}
                  borderBottom={"1px solid white"}
                  borderColor={"gray.200"}
                  height={"42px"}
                  flex={"0 0 140px"}
                  onClick={handleSelectFolderClicked}>
                  Select Folder
                </Button>
              </HStack>
            </VStack>
            {exists(sourceError) ? (
              <FormErrorMessage>{sourceError.message ?? "Error"}</FormErrorMessage>
            ) : (
              <FormHelperText></FormHelperText>
            )}
          </FormControl>

          <Box width="full" height="full" padding={4} marginBottom="8">
            <CollapsibleSection title="Or use one of our quickstart examples" showIcon isOpen>
              <FormControl isRequired isInvalid={exists(sourceError)}>
                <Wrap spacing={3} marginTop={"10px"} justify="center">
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={PythonSvg}
                      source={"https://github.com/microsoft/vscode-remote-try-python"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={NodejsPng}
                      source={"https://github.com/microsoft/vscode-remote-try-node"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={GoPng}
                      source={"https://github.com/Microsoft/vscode-remote-try-go"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={RustSvg}
                      source={"https://github.com/microsoft/vscode-remote-try-rust"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={JavaPng}
                      source={"https://github.com/microsoft/vscode-remote-try-java"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={PhpSvg}
                      source={"https://github.com/microsoft/vscode-remote-try-php"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={CppSvg}
                      source={"https://github.com/microsoft/vscode-remote-try-cpp"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                  <WrapItem padding={"1"}>
                    <ExampleCard
                      image={DotnetcorePng}
                      source={"https://github.com/microsoft/vscode-remote-try-dotnet"}
                      currentSource={currentSource}
                      setValue={setValue}
                    />
                  </WrapItem>
                </Wrap>
              </FormControl>
            </CollapsibleSection>
          </Box>
        </VStack>

        <CollapsibleSection title={"Advanced Options"} showIcon>
          <VStack spacing="10" maxWidth={"1024px"}>
            <HStack spacing="8" alignItems={"top"} width={"100%"} justifyContent={"start"}>
              <FormControl isRequired isInvalid={exists(providerError)}>
                <FormLabel>Provider</FormLabel>
                <Select {...register(FieldName.PROVIDER)}>
                  {providerOptions.map((providerID) => (
                    <option key={providerID} value={providerID}>
                      {providerID}
                    </option>
                  ))}
                </Select>
                {exists(providerError) ? (
                  <FormErrorMessage>{providerError.message ?? "Error"}</FormErrorMessage>
                ) : (
                  <FormHelperText>Use this provider to create the workspace.</FormHelperText>
                )}
              </FormControl>
              <FormControl isRequired isInvalid={exists(defaultIDEError)}>
                <FormLabel>Default IDE</FormLabel>
                <Select {...register(FieldName.DEFAULT_IDE)}>
                  {idesQuery.data?.map((ide) => (
                    <option key={ide.name} value={ide.name!}>
                      {ide.displayName}
                    </option>
                  ))}
                </Select>
                {exists(defaultIDEError) ? (
                  <FormErrorMessage>{defaultIDEError.message ?? "Error"}</FormErrorMessage>
                ) : (
                  <FormHelperText>
                    DevPod will open this workspace with the selected IDE by default. You can still
                    change your default IDE later.
                  </FormHelperText>
                )}
              </FormControl>
            </HStack>
            <HStack spacing="8" alignItems={"top"} width={"100%"} justifyContent={"start"}>
              <FormControl isInvalid={exists(idError)}>
                <FormLabel>Workspace Name</FormLabel>
                <Input
                  placeholder="my-workspace"
                  type="text"
                  {...register(FieldName.ID)}
                  onChange={(e) => {
                    setValue(FieldName.ID, e.target.value, {
                      shouldDirty: true,
                    })

                    if (/[^a-z0-9-]+/.test(e.target.value)) {
                      setError(FieldName.ID, {
                        message: "Name can only consist of lower case letters, numbers and dashes",
                      })
                    } else {
                      clearErrors(FieldName.ID)
                    }
                  }}
                />
                {exists(idError) ? (
                  <FormErrorMessage>{idError.message ?? "Error"}</FormErrorMessage>
                ) : (
                  <FormHelperText>
                    This is the workspace name DevPod will use. This is an optional field and
                    usually only needed if you have an already existing workspace with the same
                    name.
                  </FormHelperText>
                )}
              </FormControl>
              <FormControl isInvalid={exists(prebuildRepositoryError)}>
                <FormLabel>Prebuild Repository</FormLabel>
                <Input
                  placeholder="ghcr.io/my-org/my-repo"
                  type="text"
                  {...register(FieldName.PREBUILD_REPOSITORY)}
                  onChange={(e) => {
                    setValue(FieldName.PREBUILD_REPOSITORY, e.target.value, {
                      shouldDirty: true,
                    })
                  }}
                />
                {exists(prebuildRepositoryError) ? (
                  <FormErrorMessage>{prebuildRepositoryError.message ?? "Error"}</FormErrorMessage>
                ) : (
                  <FormHelperText>
                    DevPod will use this repository to find prebuilds for the given workspace.
                  </FormHelperText>
                )}
              </FormControl>
            </HStack>
          </VStack>
        </CollapsibleSection>

        <Button
          variant="primary"
          marginTop="10"
          type="submit"
          disabled={formState.isSubmitting}
          isLoading={formState.isSubmitting || isSubmitLoading}>
          Create Workspace
        </Button>
      </VStack>
    </form>
  )
}

function useCreateWorkspaceParams() {
  const [searchParams] = useSearchParams()

  return useMemo(
    () => Routes.getWorkspaceCreateParamsFromSearchParams(searchParams),
    [searchParams]
  )
}
