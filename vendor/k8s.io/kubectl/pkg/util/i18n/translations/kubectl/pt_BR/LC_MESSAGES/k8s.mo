��    f      L  �   |      �  z   �  �   	  <   �	  S   
  <   b
  c  �
  �    .   �  "   �  4   
     ?     \    {  X   �  o   �    J  v   L  t   �  �  8  ;   �  [   9  J   �  a   �  �   B  �      �   �  %   u  W   �     �  u     4   �  -   �  3   �  2        Q  *   e  .   �  *   �  0   �  0     0   L  "   }     �  *   �  A   �     +  )   I     s     �      �  (   �     �  `     �   m  �   	     �     �  $   �     �       a   0  s   �  B     +   I  +   u  6   �  q   �  /   J   1   z   '   �      �   &   �   %   !  (   :!  #   c!      �!     �!  9   �!     "      "  #   :"  �   ^"  H   �"  &   *#  e   Q#  �   �#  E   �$  a   �$  �   E%  �   &     �&     �&  =   '  $   T'     y'  &   �'  +   �'     �'  r   (     t(  /   �(  a  �(  �   *  �   �*  H   h+  W   �+  @   	,  �  J,  �  �-  .   �/  &   �/  7   0  ;   N0  ;   �0  �   �0  Z   �1  q   2  %  �2  �   �3  �   A4    �4  C   �6  �   7  j   �7  b   8  �   q8  �   :9  �   :  5   �:  k   �:  '   Z;  �   �;  8   <  1   E<  D   w<  2   �<     �<  )   =  7   5=  5   m=  3   �=  *   �=  +   >  -   .>     \>  *   x>  U   �>     �>  /   ?     D?     a?  "   ~?  0   �?     �?  e   �?  �   U@  �   A     �A  !   �A  &   �A  !   B     5B  b   PB  �   �B  G   =C  1   �C  -   �C  B   �C  w   (D  9   �D  9   �D  (   E     =E  '   WE  )   E  5   �E  3   �E  .   F     BF  G   bF  -   �F     �F  #   �F  �   G  N   �G  #   �G  c   !H  �   �H  A   yI  `   �I  �   J  �   �J     �K      L  I   L  &   fL  !   �L  &   �L  9   �L  "   M  �   3M     �M  7   �M     	   H       -   #                  3      `       d                  C          I       A          1           >   0          !           "   (       L   %       5   J   ?   4   )   b   Z   @      f   F       =         ;              c   ^         9   [   e   M      a   ,      S   '      \          Q             .   V   T   W       B      Y          E      6      :      X   &   /       P       D   K      U   2                      _   7   ]   <   8           R          $      G              O   N   *   
   +    
		  # Show metrics for all nodes
		  kubectl top node

		  # Show metrics for a given node
		  kubectl top node NODE_NAME 
		# Get the documentation of the resource and its fields
		kubectl explain pods

		# Get the documentation of a specific field of a resource
		kubectl explain pods.spec.containers 
		# Print flags inherited by all commands
		kubectl options 
		# Print the client and server versions for the current context
		kubectl version 
		# Print the supported API versions
		kubectl api-versions 
		# Show metrics for all pods in the default namespace
		kubectl top pod

		# Show metrics for all pods in the given namespace
		kubectl top pod --namespace=NAMESPACE

		# Show metrics for a given pod and its containers
		kubectl top pod POD_NAME --containers

		# Show metrics for the pods defined by label name=myLabel
		kubectl top pod -l name=myLabel 
		Convert config files between different API versions. Both YAML
		and JSON formats are accepted.

		The command takes filename, directory, or URL as input, and convert it into format
		of version specified by --output-version flag. If target version is not specified or
		not supported, convert to latest version.

		The default output will be printed to stdout in YAML format. One can use -o option
		to change to output destination. 
		Create a namespace with the specified name. 
		Create a role with single rule. 
		Create a service account with the specified name. 
		Mark node as schedulable. 
		Mark node as unschedulable. 
		Set the latest last-applied-configuration annotations by setting it to match the contents of a file.
		This results in the last-applied-configuration being updated as though 'kubectl apply -f <file>' was run,
		without updating any other parts of the object. 
	  # Create a new namespace named my-namespace
	  kubectl create namespace my-namespace 
	  # Create a new service account named my-service-account
	  kubectl create serviceaccount my-service-account 
	Create an ExternalName service with the specified name.

	ExternalName service references to an external DNS address instead of
	only pods, which will allow application authors to reference services
	that exist off platform, on other clusters, or locally. 
	Help provides help for any command in the application.
	Simply type kubectl help [path to command] for full details. 
    # Create a new LoadBalancer service named my-lbs
    kubectl create service loadbalancer my-lbs --tcp=5678:8080 
    # Dump current cluster state to stdout
    kubectl cluster-info dump

    # Dump current cluster state to /path/to/cluster-state
    kubectl cluster-info dump --output-directory=/path/to/cluster-state

    # Dump all namespaces to stdout
    kubectl cluster-info dump --all-namespaces

    # Dump a set of namespaces to /path/to/cluster-state
    kubectl cluster-info dump --namespaces default,kube-system --output-directory=/path/to/cluster-state 
    Create a LoadBalancer service with the specified name. A comma-delimited set of quota scopes that must all match each object tracked by the quota. A comma-delimited set of resource=quantity pairs that define a hard limit. A label selector to use for this budget. Only equality-based selector requirements are supported. A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.) Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP. An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field. Approve a certificate signing request Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing). Attach to a running container ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service. ClusterRole this ClusterRoleBinding should reference ClusterRole this RoleBinding should reference Convert config files between different API versions Copy files and directories to and from containers. Create a TLS secret Create a namespace with the specified name Create a secret for use with a Docker registry Create a secret using specified subcommand Create a service account with the specified name Delete the specified cluster from the kubeconfig Delete the specified context from the kubeconfig Deny a certificate signing request Describe one or many contexts Display clusters defined in the kubeconfig Display merged kubeconfig settings or a specified kubeconfig file Display one or many resources Drain node in preparation for maintenance Edit a resource on the server Email for Docker registry Execute a command in a container Forward one or more local ports to a pod Help about any command If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP' If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. Mark node as schedulable Mark node as unschedulable Mark the provided resource as paused Modify certificate resources. Modify kubeconfig files Name or number for the port on the container that the service should direct traffic to. Optional. Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used. Output shell completion code for the specified shell (bash or zsh) Password for Docker registry authentication Path to PEM encoded public key certificate. Path to private key associated with given certificate. Precondition for resource version. Requires that the current resource version match this value in order to scale. Print the client and server version information Print the list of flags inherited by all commands Print the logs for a container in a pod Resume a paused resource Role this RoleBinding should reference Run a particular image on the cluster Run a proxy to the Kubernetes API server Server location for Docker registry Set specific features on objects Set the selector on a resource Show details of a specific resource or group of resources Show the status of the rollout Synonym for --target-port The image for the container to run. The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server The minimum number or percentage of available pods this budget requires. The name for the newly created object. The name for the newly created object. If not specified, the name of the input resource will be used. The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'. The network protocol for the service to be created. Default is 'TCP'. The port that the service should serve on. Copied from the resource being exposed, if unspecified The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note that server side components may assign limits depending on the server configuration, such as limit ranges. The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges. The type of secret to create Undo a previous rollout Update resource requests/limits on objects with pod templates Update the annotations on a resource Update the labels on a resource Update the taints on one or more nodes Username for Docker registry authentication View rollout history Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory dummy restart flag) kubectl controls the Kubernetes cluster manager Project-Id-Version: 
Report-Msgid-Bugs-To: EMAIL
PO-Revision-Date: 2023-12-11 17:03+0100
Last-Translator: Carlos Panato <ctadeu@gmail.com>
Language-Team: 
Language: pt_BR
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
X-Generator: Poedit 2.4.2
Plural-Forms: nplurals=2; plural=(n > 1);
X-Poedit-KeywordsList: 
 
		  # Mostra as métricas para todos os nodes
		  kubectl top node

		  # Mostra as métricas para um node específico
		  kubectl top node NODE_NAME 
		# Mostra a documentação do recurso e seus campos
		kubectl explain pods

		# Mostra a documentação de um campo específico de um recurso
		kubectl explain pods.spec.containers 
		# Mostra as opções herdadas por todos os comandos
		kubectl options 
		# Imprime a versão do cliente e do servidor para o contexto atual
		kubectl version 
		# Mostra as versões de API suportadas
		kubectl api-versions 
		# Mostra as métricas para todos os pods no namespace default
		kubectl top pod

		# Mostra as métricas para todos os pods em um dado namespace
		kubectl top pod —namespace=NAMESPACE

		# Mostra as métricas para um dado pod e seus containers
		kubectl top pod POD_NAME —containers

		# Mostra as métricas para os pods definidos pelo label name=myLabel
		kubectl top pod -l name=myLabel 
		Convert os arquivos de configuração para diferentes versões de API. Ambos formatos YAML
	\e JSON são aceitos.

		O command recebe o nome do arquivo, diretório ou URL como entrada, e converteno formato
		para a versão especificada pelo parametro —output-version. Se a versão desejada não é especificada ou 
		não é suportada, converte para a última versã disponível.

		A saída padrão é no formato YAML. Pode ser utilizadoa opção -o
		para mudar o formato de saída. 
		Cria um namespace com um nome especificado. 
		Cria uma role com uma única regra. 
		Cria uma conta de serviço com um nome especificado. 
		Remove a restrição de execução de workloads no node. 
		Aplica a restrição de execução de workloads no node. 
		Define a annotation last-applied-configuration configurando para ser igual ao conteúdo do arquivo.
		Isto resulta no last-applied-configuration ser atualizado quando o 'kubectl apply -f <file>' executa,
		não atualizando as outras partes do objeto. 
	  # Cria um novo namespace chamado my-namespace
	  kubectl create namespace my-namespace 
	  # Cria um novo service account chamado my-service-account
	  kubectl create serviceaccount my-service-account 
	Cria um serviço do tipo ExternalName com o nome especificado.

	Serviço ExternalName referencia um endereço externo de DNS ao invés de
	apenas pods, o que permite aos desenvolvedores de aplicações referenciar serviços
	que existem fora da plataforma, em outros clusters ou localmente. 
	Help provê ajuda para qualquer comando na aplicação.
	Digite simplesmente kubectl help [caminho do comando] para detalhes completos. 
    # Cria um novo serviço do tipo LoadBalancer chamado my-lbs
    kubectl create service loadbalancer my-lbs —tcp=5678:8080 
    # Coleta o estado corrente do cluster e exibe no stdout
    kubectl cluster-info dump

    # Coleta o estado corrente do custer para /path/to/cluster-state
    kubectl cluster-info dump --output-directory=/path/to/cluster-state

    # Coleta informação de todos os namespaces para stdout
    kubectl cluster-info dump --all-namespaces

    # Coleta o conjunto especificado de namespaces para /path/to/cluster-state
    kubectl cluster-info dump --namespaces default,kube-system --output-directory=/path/to/cluster-state 
    Cria um serviço do tipo LoadBalancer com o nome especificado. Lista de valores delimitados por vírgulas para um conjunto de escopos de quota que devem corresponder para cada objeto rastreado pela quota. Lista de valores delimitados por vírgulas ajusta os pares resource=quantity que define um limite rigído. Um seletor de label a ser usado para o PDB. Apenas seletores baseado em igualdade são suportados. Um seletor de label para ser utilizado neste serviço. Apenas seletores baseado em igualdade são suportados. Se vazio (por padrão) o seletor do replication controller ou replica set será utilizado. Um IP externo adicional (não gerenciado pelo Kubernetes) para ser usado no serviço. Se este IP for roteado para um nó, o serviço pode ser acessado por este IP além de seu IP de serviço gerado. Uma substituição inline JSON para o objeto gerado. Se não estiver vazio, ele será usado para substituir o objeto gerado. Requer que o objeto forneça um campo apiVersion válido. Aprova uma solicitação de assinatura de certificado Atribuir o seu próprio ClusterIP ou configura para 'None' para um serviço 'headless' (sem loadbalancing). Se conecta a um container em execução ClusterIP que será atribuído ao serviço. Deixe vazio para auto atribuição, ou configure para 'None' para criar um serviço headless. ClusterRole que esse ClusterRoleBinding deve referenciar ClusterRole que esse RoleBinding deve referenciar Converte arquivos de configuração entre versões de API diferentes Copia arquivos e diretórios de e para containers. Cria uma secret do tipo TLS Cria a namespace com um nome especificado Cria um secret para ser utilizado com o Docker registry Cria um secret utilizando um sub-comando especificado Cria uma conta de serviço com um nome especificado Apaga o cluster especificado do kubeconfig Apaga o contexto especificado do kubeconfig Rejeita o pedido de assinatura do certificado Mostra um ou mais contextos Mostra os clusters definidos no kubeconfig Mostra a configuração do kubeconfig mescladas ou um arquivo kubeconfig especificado Mostra um ou mais recursos Drenar o node para preparação de manutenção Edita um recurso no servidor Email para o Docker registry Executa um comando em um container Encaminhar uma ou mais portas locais para um pod Ajuda sobre qualquer comando Se não vazio, configura a afinidade de sessão para o serviço; valores válidos: 'None', 'ClientIP' Se não estiver vazio, a atualização dos annotation só terá êxito se esta for a versão do recurso atual para o objeto. Válido apenas ao especificar um único recurso. Se não estiver vazio, a atualização dos labels só terá êxito se esta for a versão do recurso atual para o objeto. Válido apenas ao especificar um único recurso. Marca o node como agendável Marca o node como não agendável Marca o recurso fornecido como pausado Edita o certificado dos recursos. Edita o arquivo kubeconfig Nome ou o número da porta em um container em que o serviço deve direcionar o tráfego. Opcional. Apenas retorna os logs após uma data específica (RFC3339). Padrão para todos os logs. Apenas um since-time / since deve ser utilizado. Saída do autocomplete de shell para um Shell específico (bash ou zsh) Senha para a autenticação do registro do Docker Caminho para a chave pública em formato PEM. Caminho para a chave private associada a um certificado fornecido. Pré-condição para a versão do recurso. Requer que a versão do recurso atual corresponda a este valor para escalar. Mostra a informação de versão do cliente e do servidor Mostra a lista de opções herdadas por todos os comandos Mostra os logs de um container em um pod Retoma um recurso pausado Role que a RoleBinding deve referenciar Executa uma imagem específica no cluster Executa um proxy para o servidor de API do Kubernetes Localização do servidor para o registro do Docker Define funcionalidades específicas em objetos Define um seletor em um recurso Mostra os detalhes de um recurso específico ou de um grupo de recursos Mostra o status de uma atualização dinamica Sinônimo para —target-port A imagem para o container executar. A política de obtenção de imagens. Se deixado em branco, este valor não será especificado pelo cliente e será utilizado o padrão do servidor Um número mínimo ou porcentagem de pods disponíveis que este budget requer. O nome para o objeto recém criado. O nome para o objeto recém criado. Se não especificado, o nome do input resource será utilizado. O nome do gerador de API a ser usado. Existem 2 geradores: 'service/v1' e 'service/v2'. A única diferença entre eles é que a porta de serviço na v1 é chamada de 'default', enquanto ela é deixada sem nome na v2. O padrão é 'service/v2'. O protocolo de rede para o serviço ser criado. Padrão é 'TCP'. A porta para que o serviço possa servir. Copiado do recurso sendo exposto, se não especificado O recurso requerido para este container.  Por exemplo, 'cpu=200m,memory=512Mi'.  Observe que os componentes do lado do servidor podem atribuir limites, dependendo da configuração do servidor, como intervalos de limite. O recurso requerido de requests para este container.  Por exemplo, 'cpu=100m,memory=256Mi'.  Observe que os componentes do lado do servidor podem atribuir requests, dependendo da configuração do servidor, como intervalos de limite. O tipo de segredo para criar Desfazer o rollout anterior Atualizar os recursos de request/limites em um objeto com template de pod Atualizar as anotações de um recurso Atualizar os labels de um recurso Atualizar o taints de um ou mais nodes Nome de usuário para a autenticação no Docker registry Visualizar o histórico de rollout Onde colocar os arquivos de saída. Se vazio ou '-' usa o stdout do terminal, caso contrário, cria uma hierarquia no diretório configurado dummy restart flag) kubectl controla o gerenciador de cluster do Kubernetes 